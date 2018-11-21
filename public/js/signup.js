new Vue({
    el: '#app',
    data:{
        status:"",
        username:"",
        password:""
    },
    methods:{
        signup(){
            axios({ method: "POST", "url": "http:/api/signup", data:'{"username":"' + this.username + '", "password":"' + this.password + '"}' , withCredentials: true}).then(result => {
                this.status = result.data;
            }, error => {
                console.error(error);
            });
        }
    }
});