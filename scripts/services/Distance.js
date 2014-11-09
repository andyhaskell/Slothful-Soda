app.factory('Distance', function(){
  //Earth's radius in miles
  var RADIUS_OF_EARTH = 3959;

  //Degrees to radians function
  var radians = function(degrees){
    return degrees * (Math.PI/180);
  };

  //Great circle distance function
  var distance = function(lat1, lng1, lat2, lng2){
    var lat1  = radians(lat1),
        lng1  = radians(lng1),
        lat2  = radians(lat2),
        lng2  = radians(lng2),
        angle = Math.atan(
                  Math.sqrt(
                    Math.pow((Math.cos(lat2)*Math.sin(lng2-lng1)),2) +
                    Math.pow((Math.cos(lat1)*Math.sin(lat2)-Math.sin(lat1)*
                    Math.cos(lat2)*Math.cos(lat2-lat1)),2))/
                  (Math.sin(lat1)*Math.sin(lat2)+
                     Math.cos(lat1)*Math.cos(lat2)*Math.cos(lng2-lng1)));
                    
    return angle * RADIUS_OF_EARTH;
  };

  return {distance: distance};
});
