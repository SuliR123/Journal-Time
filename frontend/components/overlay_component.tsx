import ButtonHolder from "./button_holder"
import IconButton from "./icon_button"
import { SVGLine } from "./line"
import XIcon from "@/public/xIcon.svg"
import Checkmark from "@/public/checkmark.svg"
import { useEffect, useRef, useState } from "react"

interface BodyParam {
    title: string
    value: string
}

interface OverlayComponentParams {
    header: React.ReactNode
    subheading1: React.ReactNode
    subheading2: React.ReactNode
    body: BodyParam[]
    onClose: () => void
    onSubmit: () => void
    show?: boolean   
}

function createBodyComponent(bodyParam : BodyParam, lineWidth: number, index: number) {
    return (
        <div className="flex flex-col justify-between items-center gap-2 py-2" key={index}>
            <div className="flex w-full flex-row justify-between font-bold font-hack text-[16px]">
                <span>{bodyParam.title}</span>
                <span>{bodyParam.value}</span>
            </div>
            <SVGLine horizontal={true} strokeWidth={2} lineHeight={lineWidth}/>
        </div>
    )
}

export default function OverlayComponent({ header, subheading1, subheading2, body, onClose, onSubmit, show } : OverlayComponentParams) {
    // TODO: ADD actions like downloading the document as a pdf, sharing, etc

    const divRef = useRef<HTMLDivElement>(null);
    const [width, setWidth] = useState(0);

    useEffect(() => {
        if (divRef && divRef.current) {
            setWidth(divRef.current.offsetWidth * .95);
        }
    }, [show]); // Recalculate when show changes
    
    if(!show) {
        return null
    }

    return (
        <div className="flex items-center justify-center absolute w-[100vw] h-[100vh] z-20">
            <div className="flex items-center justify-center w-[40%] h-[70%] bg-background-color z-40 rounded-lg">
                <div ref={divRef} className="flex flex-col justify-start items-start w-[93%] h-[93%] px-4">
                    <div className="flex flex-col justify-start items-start w-full h-[25%] gap-5">
                        <div className="flex flex-col justify-center items-start w-full h-[50%]">
                            {header}
                        </div>
                        <div className="flex flex-col justify-center items-start w-full h-[40%] gap-y-2 font-hack font-bold text-[18px]">
                            {subheading1}
                            {subheading2}
                        </div>
                    </div>
                    
                    <div className="flex flex-col items-start justify-center py-4 w-full h-[75%]">
                        {body.map((item, index) => (createBodyComponent(item, width, index)))}
                    </div>

                    <div className="flex w-full flex-row justify-center items-center">
                        <ButtonHolder>
                            <IconButton icon={<XIcon />} onClick={() => {onClose()}}/>
                            <IconButton icon={<Checkmark />} onClick={() => {onSubmit()}}/>
                        </ButtonHolder>
                    </div>
                    
                </div>
            </div>
            <div className="w-full h-full absolute bg-background-color z-30 opacity-50 brightness-25" />
        </div>
    )
}