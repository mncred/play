export namespace build {
	
	export class Info {
	    time: string;
	    origin: string;
	    branch: string;
	    commit: string;
	    commitAuthor: string;
	    commitEmail: string;
	    commitTag: string;
	    compiler: string;
	    wails: string;
	    os: string;
	    arch: string;
	    logLevel: string;
	    logsDir: string;
	    appId: string;
	    windowTitle: string;
	    windowWidth: string;
	    windowHeight: string;
	    windowWidthMin: string;
	    windowWidthMax: string;
	    windowHeightMin: string;
	    windowHeightMax: string;
	    windowDisableResize: string;
	    windowFrameless: string;
	    windowDisableMaximize: string;
	    windowDisableMinimize: string;
	    windowDisableClose: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.time = source["time"];
	        this.origin = source["origin"];
	        this.branch = source["branch"];
	        this.commit = source["commit"];
	        this.commitAuthor = source["commitAuthor"];
	        this.commitEmail = source["commitEmail"];
	        this.commitTag = source["commitTag"];
	        this.compiler = source["compiler"];
	        this.wails = source["wails"];
	        this.os = source["os"];
	        this.arch = source["arch"];
	        this.logLevel = source["logLevel"];
	        this.logsDir = source["logsDir"];
	        this.appId = source["appId"];
	        this.windowTitle = source["windowTitle"];
	        this.windowWidth = source["windowWidth"];
	        this.windowHeight = source["windowHeight"];
	        this.windowWidthMin = source["windowWidthMin"];
	        this.windowWidthMax = source["windowWidthMax"];
	        this.windowHeightMin = source["windowHeightMin"];
	        this.windowHeightMax = source["windowHeightMax"];
	        this.windowDisableResize = source["windowDisableResize"];
	        this.windowFrameless = source["windowFrameless"];
	        this.windowDisableMaximize = source["windowDisableMaximize"];
	        this.windowDisableMinimize = source["windowDisableMinimize"];
	        this.windowDisableClose = source["windowDisableClose"];
	    }
	}

}

