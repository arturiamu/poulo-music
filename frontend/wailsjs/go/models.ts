export namespace models {
	
	export class VO {
	    code: number;
	    data: any;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new VO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.data = source["data"];
	        this.message = source["message"];
	    }
	}

}

