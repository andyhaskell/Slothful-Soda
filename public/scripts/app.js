app = angular.module('app', ['ngRoute']);

app.config(function($routeProvider, $locationProvider){
  $routeProvider
    .when('/', {templateUrl: '/partials/main.html'})
    .when('/map', {templateUrl: '/partials/map.html',
                   controller: 'MapCtrl'})
    .otherwise({redirectTo: '/'});

  $locationProvider.html5Mode(true);
});
