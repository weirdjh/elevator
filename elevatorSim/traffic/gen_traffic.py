import numpy as np
import sys, getopt, io
import datetime, time
import json
 
"""
Make Traffic Event which is generated by users using building

Parameters
----------
n : int
    Number of Users.
o : outputfile, optional
    By default, output file named as "user_traffic_#"

Return Example
-------
{
    "users": [
        {
            "userInfo": {
                "userID" : "1",
                "weight": 60
            },
            "moves":[
                {"at": "00:00:05.14", "from": 1, "to": 5}, 
                {"at": "00:00:08.33", "from": 5, "to": 1}
            ]
        },
    ]
}

Notes
-----
1. Users must enters the first floor of the building.
2. Users moves around the building using the elevator.
3. Users goes out through the first floor of the building.
"""

MAXFLOOR = 20
HOTTIME1 = 10
HOTTIME2 = 60
TIME_FORMAT = '%Y-%m-%d %H:%M:%S.%f'
TIME_BASE = '2018-01-01 00:00:00.00'

def argument_parsing():
    try:
        myopts, _ = getopt.getopt(sys.argv[1:],"n:o:")
    except getopt.GetoptError as e:
        print (str(e))
        print("Usage: %s -n nr_user -o output" % sys.argv[0])
        sys.exit(2)
    
    nr_user = 0
    ofile = 'user_traffic'
    for o, a in myopts:
        if o == '-n':
            nr_user = a
            ofile = ofile + "_" + nr_user
        elif o == '-o':
            ofile=a

    if nr_user == 0 :
        print("Needs number of Users as arg")
        sys.exit(2)
    
    return int(nr_user), ofile+'.json'


def next_time(time, add) :
    return time + datetime.timedelta(microseconds=add*(10**6))

def init_time() :
    base_time = datetime.datetime.strptime(TIME_BASE, TIME_FORMAT)
    #.strftime(format[:-2])
    
    elapse1 = np.random.normal(HOTTIME1, 4, 5).tolist()
    elapse2 = np.random.normal(HOTTIME2, 4, 5).tolist()
    e = np.random.choice(elapse1 + elapse2)

    return next_time(base_time, e)


def make_move() : 

    # Floor Visits rate (except first floor)
    # 50%: 1 time, 25%: 2 times, 15%: 3 times, 10%: 4 times
    nr_visit_floor = np.random.choice([1,2,3,4], 1, p=[1/2, 1/4, 3/20, 1/10])
    
    # (1 -> 16) (16 -> 3) (3 -> 12) (12 -> 1)
    move_plan = [1]
    avoid_duplicate = 1
    idx = 0
    while idx < int(*nr_visit_floor) :
        move_next = np.random.randint(1, MAXFLOOR)
        if avoid_duplicate == move_next:
            continue
        move_plan.append(move_next)
        avoid_duplicate = move_next
        idx += 1
    move_plan.append(1)
    
    # make move as json format
    moves = []
    atTime = init_time()
    fromTo = [(move_plan[pos], move_plan[pos + 1]) for pos in range(0, len(move_plan)-1)]
    for f,t in fromTo :
        move  = {
            "At" : atTime.strftime(TIME_FORMAT)[11:-4],
            "From" : f,
            "To" : t,
        }
        moves.append(move)

        # next time = elevator moving time + 10 seconds (change later)
        atTime = next_time(atTime, abs(f-t) + 10)

    return moves

def make_user (id) :
    mu, sigma = 65, 7
    weight = int(np.random.normal(mu, sigma, 1))
    
    user = {
        'userInfo' : {
            "userID" : str(id),
            "weight" : weight,
        },
        'moves' : make_move(),
    }
    return user

def write_to_jsonfile(data, ofile):
    with io.open(ofile, 'w', encoding='utf8') as outfile:
        str_ = json.dumps(data,
                        indent=4, sort_keys=True,
                        separators=(',', ': '), ensure_ascii=False)
        outfile.write(str_)

if __name__ == "__main__" :
    nr_user, ofile = argument_parsing()

    users = []
    for i in range(1, nr_user) :
        users.append(make_user(i))

    data = {'Users': users}    

    write_to_jsonfile(data, ofile)