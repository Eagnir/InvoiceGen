import Vue from "vue";

Vue.filter('igtime', function (value: any) {
  if (!value) return ''
    value = value.toString();
    // TODO: Implement time conversion from UTC to local and display string in format hh:mm am/pm
  return value;
})