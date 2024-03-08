export namespace model {
	
	export class DT {
	    id: number;
	    ampId: number;
	    class: string;
	    mode: string;
	    topology: string;
	
	    static createFrom(source: any = {}) {
	        return new DT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.ampId = source["ampId"];
	        this.class = source["class"];
	        this.mode = source["mode"];
	        this.topology = source["topology"];
	    }
	}
	export class Parameter {
	    id: number;
	    name: string;
	    type: string;
	    value: string;
	    valueNumber: number;
	    allowedValue: string[];
	    min: number;
	    max: number;
	
	    static createFrom(source: any = {}) {
	        return new Parameter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.value = source["value"];
	        this.valueNumber = source["valueNumber"];
	        this.allowedValue = source["allowedValue"];
	        this.min = source["min"];
	        this.max = source["max"];
	    }
	}
	export class PedalBoardItem {
	    active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PedalBoardItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.active = source["active"];
	    }
	}
	export class Preset {
	    id: number;
	    idStr: string;
	    dts: DT[];
	    items: PedalBoardItem[];
	    name: string;
	    parameters: Parameter[];
	    setId: number;
	
	    static createFrom(source: any = {}) {
	        return new Preset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.idStr = source["idStr"];
	        this.dts = this.convertValues(source["dts"], DT);
	        this.items = this.convertValues(source["items"], PedalBoardItem);
	        this.name = source["name"];
	        this.parameters = this.convertValues(source["parameters"], Parameter);
	        this.setId = source["setId"];
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
	export class Set {
	    id: number;
	    name: string;
	    presets: Preset[];
	
	    static createFrom(source: any = {}) {
	        return new Set(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.presets = this.convertValues(source["presets"], Preset);
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
	export class Pod {
	    currentPresetId?: number;
	    currentSetId?: number;
	    sets: Set[];
	
	    static createFrom(source: any = {}) {
	        return new Pod(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.currentPresetId = source["currentPresetId"];
	        this.currentSetId = source["currentSetId"];
	        this.sets = this.convertValues(source["sets"], Set);
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

