new Vue({
    el: '#train',
    data:{
        status:"",
        speed:0,
        direction:0
    },
    methods:{
        submitTrainCommand(){
            console.log("asdfasdf");
            axios({ method: "POST", "url": "http:/api/settraincommand", data:'{"speed":' + this.speed + ', "direction":' + this.direction + '}' , withCredentials: true}).then(result => {
                this.status = result.data;
            }, error => {
                console.error(error);
            });

        }
    }
});
