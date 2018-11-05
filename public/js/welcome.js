
new Vue({
    el: '#railway',
    data:{
        jsonInfo:"a",
        speed:6,
        direction:0,
        firstswitch:1,
        secondswitch:0
    },
    mounted(){
        axios({ method: "GET", "url": "http:/api/railwayinfo" }).then(result => {
            this.direction = result.data["train"]["direction"];
            this.speed = result.data["train"]["speed"];
            this.firstswitch = result.data["railway"]["firstswitch"];
            this.secondswitch = result.data["railway"]["secondswitch"];
        }, error => {
            console.error(error);
        });
    },
    methods:{
        getInfo() {
            axios({ method: "GET", "url": "http:/api/railwayinfo" }).then(result => {
                this.direction = result.data["train"]["direction"];
                this.speed = result.data["train"]["speed"];
                this.firstswitch = result.data["railway"]["firstswitch"];
                this.secondswitch = result.data["railway"]["secondswitch"];
            }, error => {
                console.error(error);
            });
        }

    }
});