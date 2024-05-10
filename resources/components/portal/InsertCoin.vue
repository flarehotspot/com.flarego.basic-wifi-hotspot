<template>
    <p v-if="loading">Loading...</p>
</template>

<script>
define(function () {
    return {
        template: template,
        data: function(){
            return {
                loading: true,
                error: null
            }
        },
        mounted: function () {
            var self = this;
            $flare.http.get('<% .Helpers.UrlForRoute "portal.purchase.wifi" %>').catch(function(err) {
                console.log(err)
                self.error = err;
            }).finally(function(){
                self.loading = false;
            });
        }
    };
});
</script>
