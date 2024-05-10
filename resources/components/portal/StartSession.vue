<template>
    <div>
        <p>Starting session...</p>
    </div>
</template>

<script>
    define(function() {
        return {
            template: template,
            mounted: function() {
                $flare.http.post('<% .Helpers.UrlForRoute "portal.sessions.start" %>')
            }
        }
    })
</script>
