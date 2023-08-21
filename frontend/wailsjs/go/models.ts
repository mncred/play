export namespace config {
	
	export class WindowAppearanceBackground {
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new WindowAppearanceBackground(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.color = source["color"];
	    }
	}
	export class WindowAppearance {
	    roundness: number;
	    background: WindowAppearanceBackground;
	
	    static createFrom(source: any = {}) {
	        return new WindowAppearance(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.roundness = source["roundness"];
	        this.background = this.convertValues(source["background"], WindowAppearanceBackground);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class WindowSize {
	    default: number;
	    min: number;
	    max: number;
	
	    static createFrom(source: any = {}) {
	        return new WindowSize(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.default = source["default"];
	        this.min = source["min"];
	        this.max = source["max"];
	    }
	}
	export class WindowHeaderText {
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new WindowHeaderText(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.color = source["color"];
	    }
	}
	export class WindowHeaderBackground {
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new WindowHeaderBackground(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.color = source["color"];
	    }
	}
	export class WindowHeader {
	    native: boolean;
	    title: string;
	    maximise: boolean;
	    background: WindowHeaderBackground;
	    text: WindowHeaderText;
	
	    static createFrom(source: any = {}) {
	        return new WindowHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.native = source["native"];
	        this.title = source["title"];
	        this.maximise = source["maximise"];
	        this.background = this.convertValues(source["background"], WindowHeaderBackground);
	        this.text = this.convertValues(source["text"], WindowHeaderText);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Window {
	    header: WindowHeader;
	    width: WindowSize;
	    height: WindowSize;
	    resizable: boolean;
	    appearance: WindowAppearance;
	
	    static createFrom(source: any = {}) {
	        return new Window(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.header = this.convertValues(source["header"], WindowHeader);
	        this.width = this.convertValues(source["width"], WindowSize);
	        this.height = this.convertValues(source["height"], WindowSize);
	        this.resizable = source["resizable"];
	        this.appearance = this.convertValues(source["appearance"], WindowAppearance);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Log {
	    level: string;
	    output: string;
	
	    static createFrom(source: any = {}) {
	        return new Log(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.level = source["level"];
	        this.output = source["output"];
	    }
	}
	export class Debug {
	    inspector: boolean;
	    log: Log;
	
	    static createFrom(source: any = {}) {
	        return new Debug(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.inspector = source["inspector"];
	        this.log = this.convertValues(source["log"], Log);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Config {
	    id: string;
	    debug: Debug;
	    window: Window;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.debug = this.convertValues(source["debug"], Debug);
	        this.window = this.convertValues(source["window"], Window);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	
	
	
	

}

