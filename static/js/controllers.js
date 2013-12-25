'use strict'

// Controllers

function FrontCtrl ($scope, $http, $state, $rootScope) {
	$http.get('/bloginfo').success(function (data) {
		console.log(data)
		$rootScope.title = $scope.title = data.title;
		$rootScope.subtitle = $scope.subtitle = data.subtitle;
		$scope.avatarlink = 'http://www.gravatar.com/avatar/' + data.avatarhash + '?s=100';
		$rootScope.blog = data
	});
}

function ArticleListCtrl ($scope, $http) {
	$http.get('/articles', {params: {status: "published"}}).success(function (data) {
		console.log(data);
		for (var i in data) {
			data[i].date = moment(data[i].date);
		}
		$scope.articles = data;
	});
}

function ArticleDetailCtrl ($scope, $http, $stateParams) {
	console.log('ArticleDetailCtrl');
	$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
		$scope.article = data;
	});
}

function AdminCtrl ($scope, $http, $state, $rootScope) {
	$http.get('/admin/user').success(function (data) {
		if (data.code == 404) {
			$state.transitionTo('admin.login');
		} else if (data.code == 200) {
			$rootScope.user = data.data;
		}
	});
}

function DashboardCtrl ($scope, $http, $state, $rootScope) {
	console.log($rootScope.user);
	if ($rootScope.user == undefined) {
		$state.transitionTo('admin.login');
	}
}

function AdminArticleListCtrl ($scope, $http, $notification, $state, $rootScope) {
	if ($rootScope.user == undefined) {
		$state.transitionTo('admin.login');
	}

	$http.get('/articles').success(function (data) {
		$scope.articles = data;
	});

	$scope.delete = function (id) {
		$http.delete('/admin/post/delete', {params: {id: id}}).success(function (data, status) {
			if (data.code == 200) {
				for (var i = 0; i < $scope.articles.length; i++) {
					var a = $scope.articles[i];
					if (a.id == id) {
						$scope.articles.splice(i, 1);
						$notification.success('Success', 'Delete "' + a.title + '" succeed.');
						break;
					}
				}
			}
		});
	};
}

function AdminArticleEditCtrl ($scope, $http, $stateParams, $state, $notification, $rootScope) {
	if ($rootScope.user == undefined) {
		$state.transitionTo('admin.login');
	}

	var postUrl = '/admin/post';

	// Check state between admin.new and admin.article
	var articleNew = $stateParams.articleId == undefined;

	if (!articleNew) {
		// console.log($stateParams.articleId);
		$http.get('/article', {params: {id: $stateParams.articleId}}).success(function (data) {
			$scope.article = data
			$scope.date = moment(data.date).format('YYYY-MM-DD');
		});

		postUrl = postUrl + '/update';
	} else {
		postUrl = postUrl + '/new';
	}

	// Save draft form button action
	$scope.draft = function () {
		var article = $scope.article;
		article.status = 'draft';
		
		article.date = generateDate($scope.date).format();
		if (article.date != undefined) {
			article.author = 'DYZ';
			console.log(article);
			$http.post(postUrl, article).success(function (data, status) {
				console.log(data);
				if (data.code == 200) {
					$state.transitionTo('admin.articles');
					$notification.success('Success', 'Save draft succeed.');
				} else{
					$notification.error('Error', 'Something is wrong.');
				}
			})
		} else {

		}
	};

	// Post form button action
	$scope.post = function () {
		var article = $scope.article;
		article.status = 'published';
		article.date = generateDate($scope.date).format();

		if (article.date != undefined) {
			article.author = 'DYZ';
			console.log(article);
			$http.post(postUrl, article).success(function (data, status) {
				console.log(data);
				if (data.code == 200) {
					$state.transitionTo('admin.articles');
					$notification.success('Success', 'Post succeed.');
				} else{
					$notification.error('Error', 'Something is wrong.');
				}
			})
		} else {

		}
	};

	var generateDate = function (d) {
		var date = moment();
		if (d == undefined || d == '') {
			return date;
		}

		var temp = moment(d, 'YYYY-MM-DD', true);
		if (!temp.isValid()) {
			return undefined;
		}

		date.date(temp.date());
		date.month(temp.month());
		date.year(temp.year());

		return date;
	};
}

function AdminLoginCtrl ($scope, $http, $state, $rootScope) {
	$scope.login = function () {
		var user = {email: $scope.user.email, password: $scope.user.password};
		$http.post('/admin/user/login', user).success(function (data) {
			console.log(data);
			if (data.code == 200) {
				$rootScope.user = data.data;
				$state.transitionTo('admin.dashboard');
			};
		});
	}
}