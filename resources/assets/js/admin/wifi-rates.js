$(document).ready(function () {
  var newRateBtn = $("#new-rate-btn");
  var newRateForm = $("#new-rate-form");
  var newRateCheckbox = $("#new-rate-checkbox");
  var cancelNewRateBtn = $("#cancel-new-rate");

  newRateBtn.click(function () {
    newRateForm.show();
    newRateCheckbox.attr("checked", "checked");
  });

  cancelNewRateBtn.click(function () {
    newRateForm.hide();
    newRateCheckbox.removeAttr("checked");
  });
});
