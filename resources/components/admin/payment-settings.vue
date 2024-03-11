<template>
    <div>
        <p v-if="!flareView.data.length">No data.</p>
        <form @submit.prevent="submit">
            <div v-for="s in flareView.data">
                Amount: <input type="number" v-model="s.amount" step="0.01" min="0" required>
                Time Mins: <input type="number" v-model="s.time_mins" step="0.01" min="0" required>
                Data Mbytes: <input type="number" v-model="s.data_mb" step="0.01" min="0" required>
                <button type="button" @click="deleteEntry(s.amount)">Delete Denomination</button>
            </div>
            <button type="button" @click="addEntry">Add Denomination</button>
            
            <br>
            <button type="submit">Submit</button>
        </form>
    </div>
</template>
<script>
define(function () {
    return {
        props: ['flareView'],
        template: template,
        methods: {
            addEntry: function () {
                this.flareView.data.push({
                    amount: 0.0,
                    time_mins: 0,
                    data_mb: 0
                })
            },
            deleteEntry:function(denom) {
                var index=this.flareView.data.findIndex(function (item){
                    return item.amount===denom
                });
                if(index!==-1){
                    this.flareView.denom=denom;
                    this.flareView.data.splice(index,1);
                }
            },
            submit: function () {
                var data = this.flareView.data;
                for (var i = 0; i < data.length; i++) {
                    data[i] = {
                        amount: data[i].amount * 1,
                        time_mins: data[i].time_mins * 1,
                        data_mb: data[i].data_mb * 1
                    };
                }
                window.$flare.http.post('<% .Helpers.UrlForRoute "admin.payment-settings.save" %>', data).catch(function (err) {
                    console.log(err)
                })
            }
        }
    }
})
</script>
