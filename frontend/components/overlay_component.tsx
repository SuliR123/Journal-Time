import IconButton from "./icon_button"
import CheckmarkIcon from "@/public/checkmark.svg"
import XIcon from "@/public/xIcon.svg"
import { SVGLine } from "./line"
import { Dispatch, SetStateAction, useEffect, useRef, useState } from "react"
import ButtonHolder from "./button_holder"

interface OverlayComponentParams {
    show?: boolean
    left: React.ReactNode
    right: React.ReactNode
    toggle: Dispatch<SetStateAction<boolean>>
    onSubmit?: () => void
    onClose?: () => void
}


export default function OverlayComponent({ show = false, left, right, toggle, onSubmit, onClose} : OverlayComponentParams) {
    const divRef = useRef<HTMLDivElement>(null);
    const [height, setHeight] = useState(100);

    useEffect(() => {
        if (divRef && divRef.current) {
            setHeight(divRef.current.offsetHeight * .95);
        }
    }, [show]); // Recalculate when show changes

    if(!show) { return null }

    return (
        <div className="flex items-center justify-center absolute w-[100vw] h-[100vh] z-10">
            <div className="flex items-center justify-center z-30 w-[60vw] h-[70vh] absolute bg-background-color">
                <div className="flex flex-col items-center justify-between w-[95%] h-[95%]">
                    <div className="flex flex-row items-center justify-between w-[100%] h-[100%]">
                        <div ref={divRef} className="w-[45%] h-[90%]">
                            {left}
                        </div>
                        <SVGLine lineHeight={height} horizontal={false} strokeWidth={4}/>
                        <div className="w-[45%] h-[90%]">
                            {right}
                        </div>
                    </div>
                    <ButtonHolder>
                        <IconButton icon={<XIcon />} onClick={() => 
                            {
                                toggle(false)
                                if(onClose) {
                                    onClose()
                                }
                            }}/>
                        <IconButton icon={<CheckmarkIcon />} onClick={() => 
                        {
                            toggle(false)
                            if(onSubmit) {
                                onSubmit()
                            }
                        }}/>
                    </ButtonHolder>
                    
                </div>
            </div>
            <div className="w-full h-full absolute bg-background-color z-20 opacity-50 brightness-25" />
        </div>
    )
}