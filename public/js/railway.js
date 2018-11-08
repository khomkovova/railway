new Vue({
    el: '#railway',
    data:{
        status:"",
        firstswitch:0,
        secondswitch:0
    },
    methods:{
        submitRailwaynCommand(){
            axios({ method: "POST", "url": "http:/api/setrailwaycommand", data:'{"firstswitch":' + this.firstswitch + ', "secondswitch":' + this.secondswitch + '}' , withCredentials: true}).then(result => {
                this.status = result.data;
            }, error => {
                console.error(error);
            });

        }
    }
});