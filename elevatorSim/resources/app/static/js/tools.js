function env_input(){
    var building_max = document.getElementById("building_max").value;
    var elevator_nr = document.getElementById("elevator_nr").value; 
    gui_init(building_max, elevator_nr)
}

function gui_init(b, e) {

    // init select -> one person add
    [...document.getElementsByClassName("one_person_add")].forEach(x=>{
        [...Array(Number(b)).keys()].forEach(y=>{
            var option = document.createElement("option");
            option.text = y+1;
            x.add(option);
        });
    });

    // elevator add
    [...Array(Number(e)).keys()].forEach(i=>{
        var Div = document.createElement("div");
        Div.className = 'elevator_box';
        document.getElementById("panel").appendChild(Div);

        // align by 4
        if ( (i+1) % 5 == 0 ){
            console.log(i);
            var div_align = document.createElement("div");
            div_align.style.clear="both";
            document.getElementById("panel").appendChild(div_align);
        }
    });

    // add event handler at traffic div
    var one_traffic = document.getElementById("one_traffic");
    one_traffic.addEventListener("click", function(event){
        alert('Hi!');
    });

    var many_traffic = document.getElementById("many_traffic");
    many_traffic.addEventListener("click", function(event){
        alert('Hi!');
    });

    // clock tool event handler
    var pause = document.getElementById("pause");
    many_traffic.addEventListener("click", function(event){
        alert('Hi!');
    });

    var start = document.getElementById("start");
    many_traffic.addEventListener("click", function(event){
        alert('Hi!');
    });

    var faster = document.getElementById("faster");
    many_traffic.addEventListener("click", function(event){
        alert('Hi!');
    });
}

