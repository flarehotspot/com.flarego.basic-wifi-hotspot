<template>
    <div>
        <p v-if="!data.length">No data.</p>
        <form @submit.prevent="submit">
            <div v-for="s in data">
                Amount: <input type="number" v-model="s.amount" step="0.01" min="0" required>
                Time Mins: <input type="number" v-model="s.time_mins" step="1" min="0" required>
                Data Mbytes: <input type="number" v-model="s.data_mb" step="1" min="0" required>
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
        template: template,
        data: function(){
            return {
                data: [],
                denom: null
            }
        },
        mounted: function(){
            var self = this;
            $flare.http.get('<% .Helpers.UrlForRoute "admin.payment-settings.get" %>').then(function(data) {
                console.log(data)
                self.data = data;
            })
        },
        methods: {
            addEntry: function () {
                this.data.push({
                    amount: 0.0,
                    time_mins: 0,
                    data_mb: 0
                })
            },
            deleteEntry: function (denom) {
                var index = -1;
                for (var i = 0; i < this.data.length; i++) {
                    if (this.data[i].amount === denom) {
                        index = i;
                        break;
                    }
                }
                if (index !== -1) {
                    this.denom = denom;
                    this.data.splice(index, 1);
                }
            },
            submit: function () {
                var data = this.data;
                for (var i = 0; i < data.length; i++) {
                    data[i] = {
                        amount: data[i].amount * 1,
                        time_mins: data[i].time_mins * 1,
                        data_mb: data[i].data_mb * 1
                    };
                }
                $flare.http.post('<% .Helpers.UrlForRoute "admin.payment-settings.save" %>', data).catch(function (err) {
                    console.error(err)
                })
            }
        }
    }
})
</script>
