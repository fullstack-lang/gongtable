// insertion point for imports
import { SVGIconDB } from './svgicon-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ButtonDB {

	static GONGSTRUCT_NAME = "Button"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Icon: string = ""

	// insertion point for pointers and slices of pointers declarations
	SVGIcon?: SVGIconDB


	ButtonPointersEncoding: ButtonPointersEncoding = new ButtonPointersEncoding
}

export class ButtonPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	SVGIconID: NullInt64 = new NullInt64 // if pointer is null, SVGIcon.ID = 0

}
