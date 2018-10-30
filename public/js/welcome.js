
new Vue({
    el: '#railway',
    data:{
        jsonInfo:"a",
        speed:0,
        direction:0,
        firstswitch:0,
        secondswitch:0
    },
    // mounted() {
    //
    // },
    methods:{

        getInfo() {
            axios({ method: "GET", "url": "http://127.0.0.1:12345/api/railwayinfo" }).then(result => {
                // this.jsonInfo = result.data;
                // var obj = JSON.parse(result.data);
                // console.error(result.data);
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