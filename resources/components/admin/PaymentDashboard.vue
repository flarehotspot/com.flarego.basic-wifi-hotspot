<template>
  <b-container fluid>
    <b-card bg-variant="dark" text-variant="white" header="Payment Dashboard" class="mt-5 mt-sm-5 text-center">
      <b-form-group
        label-cols-lg="3"
        label="Price Plans"
        label-size="lg"
        label-class="font-weight-bold pt-0"
        class="mb-0"
      >
        <div v-for="(plan) in plans" :key="plan.price">
          <b-form-group
            :label="`₱${plan.price}:`"
            label-cols-sm="3"
            label-align-sm="right"
          >
            <label>Data Allocation: {{ plan.dataAllocation }} MB</label>
            <b-form-spinbutton
              class="text-center"
              :id="`sb-wrap-${plan.price}`"
              wrap
              min="0"
              max="1024"
              placeholder="in megabytes"
              size="lg"
              step="100"
              v-model="plan.dataAllocation"
            ></b-form-spinbutton>
            <label for="sb-wrap" class="mt-2">Time Allotment: {{ plan.timeInMins }}</label>
            <b-form-input
              class="mr-2 ml-2"
              :id="`range-${plan.price}`"
              v-model="plan.timeInMins"
              type="range"
              min="0"
              max="360"
              step="5"
            ></b-form-input>
          </b-form-group>
        </div>
      </b-form-group>
    </b-card>
    <b-button block squared variant="primary" @click="submitForm" class="mt-3" size="lg">Submit</b-button>
  </b-container>
  
</template>
  
<script>
define(function() {
  return {
    template: template,
    data: function() {
      return {
        plans: [
          { price: 1, dataAllocation: 0, timeInMins:'0' },
          { price: 5, dataAllocation: 0, timeInMins: '0' },
          { price: 10, dataAllocation: 0, timeInMins: '0'},
        ],  
      };
    },
    // computed: {
    //   timeLabels() {
    //     return this.plans.map(plan => {
    //       if(plan.timeInMins === 0 || plan.timeInMins === 1) {
    //         return '1 minute';
    //       } else if(plan.timeInMins % 60 === 0) {
    //         const hours = plan.timeInMins / 60;
    //         return hours === 1 ? '1 hour' : `${hours} hours`;
    //       } else {
    //         const hours = Math.floor(plan.timeInMins / 60);
    //         const minutes = plan.timeInMins % 60;
    //         if(hours === 0) {
    //           return `${minutes} minutes`;
    //         } else {
    //           return `${hours} hours ${minutes} minutes`;
    //         }
    //       }
    //     });
    //   }
    // },
    methods: {
      submitForm() {
        // Form data
        const formData = 
      this.plans.map(plan => ({
        price: plan.price,
        dataAlloc: plan.dataAllocation,
        timeInMins: plan.timeInMins*1
      }));

        window.$flare.http.post('<% .Helpers.UrlForRoute "save-settings" %>', formData)
        .then(response=>{
          console.log("Saved response:", response)
        })
        .catch(error=>{
          console.error("Error responding:",error)
        })
      } 
    }
  };
});
</script>