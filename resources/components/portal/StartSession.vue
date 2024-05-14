<template>
    <div>
        <p v-if="starting && !redirect_url">Starting session...</p>
        <p v-if="redirect_url">Redirecting to {{ redirect_url }}...</p>
    </div>
</template>

<script>
    define(function() {
        return {
            template: template,
            data: function(){
                return {
                    starting: true,
                    redirect_url: "",
                };
            },
            mounted: function() {
                var self = this;
                $flare.http.post('<% .Helpers.UrlForRoute "portal.sessions.start" %>')
                    .then(function(data){
                        self.starting = false;
                        self.redirect_url = data.redirect_url;
                        setTimeout(function() {
                            window.location.href = data.redirect_url;
                        }, 1500);
                    }).catch(function(error){
                        self.$router.push("/");
                    });
            }
        }
    })
</script>
