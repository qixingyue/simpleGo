define(function(require,exports,module){

		var queryString = "";
		var queryArr = [];
		var routeObjs = [];
		
		exports.setObjs = function(allObjs){
			routeObjs = allObjs;
		};

		exports.router = function(){
			queryString = location.search.substr(1);
			var andSplitArr = queryString.split("&");
			for(var m = 0, l = andSplitArr.length ; m < l ; m++ ){
				var kvArr = andSplitArr[m].split("=");
				queryArr[kvArr[0]] = kvArr[1];
			}
			var currentaction = "./" + queryArr["action"];
			require.async([currentaction],function(h){
				h.doAction();
			});
		};

});
