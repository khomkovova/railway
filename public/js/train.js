new Vue({
    el: '#train',
    data:{
        speed:0,
        direction:0
    },
    methods:{
        submitTrainCommand(){
            console.log("asdfasdf");
            axios({ method: "POST", "url": "http://127.0.0.1:12345/api/settraincommand", data:'{"speed":' + this.speed + ', "direction":' + this.direction + '}' , withCredentials: true}).then(result => {

                console.error(result.data);
            }, error => {
                console.error(error);
            });

        }
    }
});