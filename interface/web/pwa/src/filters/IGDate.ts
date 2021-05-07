import Vue from "vue";

Vue.filter('igdate', function (value: any) {
  if (!value) return ''
    value = value.toString();
    // TODO: Implement date conversion from UTC to local and display string in format DD MMM YYYY
  return value;
})