define(function(require,exports,module){
	
		var r = require("./router");

		var init = function(){
			var objs = ["queue"];
			r.setObjs(objs);
			r.router();
		};

		init();
});
