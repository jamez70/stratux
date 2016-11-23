ifeq ($(statuxBuild),)
stratuxBuild=`git log -n 1 --pretty=%H`
buildtime=`date +%s`
endif


WWWDIR=/var/www
all:
	mkdir -p $(WWWDIR)
	mkdir -p $(WWWDIR)/css
	cp css/*.css $(WWWDIR)/css
	mkdir -p $(WWWDIR)/js
	cp js/main.js js/addtohomescreen.min.js js/j3di-all.min.js $(WWWDIR)/js
	mkdir -p $(WWWDIR)/img
	cp img/logo*.png img/screen*.png img/world.png $(WWWDIR)/img
	mkdir -p $(WWWDIR)/maui
	mkdir -p $(WWWDIR)/maui/js
	cp maui/js/angular-ui-router.min.js maui/js/mobile-angular-ui.min.js maui/js/angular.min.js maui/js/mobile-angular-ui.gestures.min.js maui/js/mobile-angular-ui.core.min.js $(WWWDIR)/maui/js
	mkdir -p $(WWWDIR)/maui/css
	cp maui/css/mobile-angular-ui-hover.min.css maui/css/mobile-angular-ui-desktop.min.css maui/css/mobile-angular-ui-base.min.css $(WWWDIR)/maui/css
	mkdir -p $(WWWDIR)/maui/fonts
	cp maui/fonts/fontawesome-webfont.woff $(WWWDIR)/maui/fonts
	mkdir -p $(WWWDIR)/plates
	cp plates/*.html $(WWWDIR)/plates
	mkdir -p $(WWWDIR)/plates/js
	cp plates/js/*.js $(WWWDIR)/plates/js
	cp index.html stratux.appcache $(WWWDIR)
	# Mark the manifest with the git hash.
	echo "# build time: " ${buildtime} >>$(WWWDIR)/stratux.appcache
	echo "# Stratux build: " ${stratuxBuild} >>$(WWWDIR)/stratux.appcache
	
