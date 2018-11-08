
new Vue({
    el: '#railway',
    data:{
        jsonInfo:"a",
        speed:0,
        direction:0,
        firstswitch:"",
        secondswitch:0
    },
    mounted(){
        axios({ method: "GET", "url": "http:/api/railwayinfo" }).then(result => {
            this.direction = result.data["train"]["direction"];
            if (result.data["train"]["direction"] == 0){
                this.direction = "Back";
            }
            if (result.data["train"]["direction"] == 1){
                this.direction = "Forward";
            }
            this.speed = result.data["train"]["speed"];
            if (result.data["railway"]["firstswitch"] == 0){
                this.firstswitch = "Left";
            }
            if (result.data["railway"]["firstswitch"] == 1){
                this.firstswitch = "Right";
            }
            if (result.data["railway"]["firstswitch"] == 0){
                this.secondswitch = "Left";
            }
            if (result.data["railway"]["firstswitch"] == 1){
                this.secondswitch = "Right";
            }
        }, error => {
            console.error(error);
        });
    },
    methods:{
        getInfo() {
            axios({ method: "GET", "url": "http:/api/railwayinfo" }).then(result => {
                this.direction = result.data["train"]["direction"];
                if (result.data["train"]["direction"] == 0){
                    this.direction = "Back";
                }
                if (result.data["train"]["direction"] == 1){
                    this.direction = "Forward";
                }
                this.speed = result.data["train"]["speed"];
                if (result.data["railway"]["firstswitch"] == 0){
                    this.firstswitch = "Left";
                }
                if (result.data["railway"]["firstswitch"] == 1){
                    this.firstswitch = "Right";
                }
                if (result.data["railway"]["firstswitch"] == 0){
                    this.secondswitch = "Left";
                }
                if (result.data["railway"]["firstswitch"] == 1){
                    this.secondswitch = "Right";
                }
            }, error => {
                console.error(error);
            });
        }

    }
});