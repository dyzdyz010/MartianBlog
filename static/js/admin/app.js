'use strict'

angular.module('MartianBlog', ['markdownFilters']).
  config(['$routeProvider', function($routeProvider) {
  $routeProvider.
      when('/new-article', {templateUrl: 'static/views/admin/new-article.html',   controller: NewArticleCtrl}).
      otherwise({redirectTo: '/new-article'});
}]);
