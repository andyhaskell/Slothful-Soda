app = angular.module('app', ['ngRoute']);

app.config(function($routeProvider, $locationProvider){
  $routeProvider
    .when('/', {templateUrl: 'http://andyhaskell.github.io/Slothful-Soda/partials/main.html'})
    .when('/map', {templateUrl: 'http://andyhaskell.github.io/Slothful-Soda/partials/map.html',
                   controller: 'MapCtrl'})
    .otherwise({redirectTo: '/'});

  $locationProvider.html5Mode(true);
});
