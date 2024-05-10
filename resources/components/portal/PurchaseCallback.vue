<template>
    <div>
        <h3>Thank You !!!</h3>
        <router-link to="/"><button>Go to Home</button></router-link>
    </div>
</template>

<script>
define(function () {
    return {
        template: template,
        mounted: function() {
            $flare.http.get('<% .Helpers.UrlForRoute "portal.purchase.callback" %>').catch(function(err) {
                console.log(err);
            })
        }
    };
});
</script>
