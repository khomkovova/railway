new Vue({
    el: '#railway',
    data:{
        firstswitch:0,
        secondswitch:0
    },
    methods:{
        submitRailwaynCommand(){
            console.log("asdfasdf");
            axios({ method: "POST", "url": "http://127.0.0.1:12345/api/setrailwaycommand", data:'{"firstswitch":' + this.firstswitch + ', "secondswitch":' + this.secondswitch + '}' , withCredentials: true}).then(result => {

                console.error(result.data);
            }, error => {
                console.error(error);
            });

        }
    }
});