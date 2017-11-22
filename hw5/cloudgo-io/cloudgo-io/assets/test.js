$(document).ready(function() {
    setInterval(function() {
      var dateNtime = new Date();
      var year = dateNtime.getFullYear();
      var month = dateNtime.getMonth() + 1;
      var day = dateNtime.getDate();
      var dateStr = month + "/" + day + ", " + year;
      $('.index-date').html("Today is " + dateStr);
    }, 1000)
})
