app.controller('MapCtrl', function(Distance, $scope){
  $scope.locations     = [{"Id":1,"Name":"Tufts Res Quad","Lat":42.408565,"Lng":-71.121765},{"Id":2,"Name":"Hodgkins Park","Lat":42.399566,"Lng":-71.124595},{"Id":3,"Name":"Danehy Park","Lat":42.388306,"Lng":-71.137507},{"Id":4,"Name":"Lederman Park","Lat":42.363649,"Lng":-71.071774}];
  $scope.markers       = [];
  $scope.lastDistance  = 5;
  $scope.distance      = 5;
  $scope.map           = null;

  var ourLat = 42.388282,
      ourLng = -71.153968;

  var initMap = function(){
    var options = {
      center   : {lat: ourLat, lng: ourLng},
      mapTypeId: google.maps.MapTypeId.HYBRID,
      zoom     : 12
    };

    var map = new google.maps.Map(document.getElementById('map'), options);

    for (var i = 0; i < $scope.locations.length; i++) {
        var currentLoc = $scope.locations[i];
        var markerOptions = {
            position: new google.maps.LatLng(currentLoc.Lat, currentLoc.Lng),
            title   : currentLoc.Name,
            visible : true,
            map     : map
        };
        var marker = new google.maps.Marker(markerOptions);
        $scope.markers.push(marker);
    }
    $scope.refreshDistance();
  };

  //Initialize the map and the markers
  $scope.initialize = function(){
    initMap();
  };

  //Show only markers on the map within the specified number of miles
  $scope.refreshDistance = function(){
    var miles = parseFloat($scope.distance);
    miles = isNaN(miles) ? $scope.lastDistance : miles;
    $scope.lastDistance = miles;

    var markers = $scope.markers;

    for (var i = 0; i < markers.length; i++) {
      var pos = markers[i].getPosition();
      var lat = pos.lat(),
          lng = pos.lng();

      markers[i].setVisible(Distance.distance(ourLat,ourLng,lat,lng) <= miles);
    }
  }
});
