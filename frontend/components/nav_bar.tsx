'use client';
import ClockIcon from "@/public/clock.svg"
import NotebookIcon from "@/public/notebook.svg"
import NoteIcon from "@/public/note.svg"
import StatsIcon from "@/public/stats.svg"
import ProfileIcon from "@/public/profile.svg"
import IconLink from "./icon_link";
import Logo from "@/public/logo.svg"
import { SVGLine } from "./line";
import { useEffect, useRef, useState } from "react";

export default function NavBar() {
    const divRef = useRef<HTMLDivElement>(null);
    const [height, setHeight] = useState(0)
    const [hover, setHover] = useState(false)

    useEffect(() => {
        if(divRef.current) {
            setHeight(divRef.current.offsetHeight)
        }
    }, []); // Recalculate when show changes

    return (
    <div ref={divRef} className="z-10 absolute w-full h-full pointer-events-none">
        <div className="flex flex-row items-start justify-between w-full h-full">
            <div className="flex flex-row h-full pointer-events-auto"
                onMouseEnter={() => setHover(true)}
                onMouseLeave={() => setHover(false)}
                >
                <div className="flex flex-col items-center w-[5vw] h-full">
                    <div className="flex flex-col items-center justify-between gap-10 h-[95%] z-20 pt-7">
                        <IconLink icon={<Logo />} link="/" size="lg"/>
                        {hover && <div className="flex flex-col items-center gap-10 w-full h-full">
                            <IconLink icon={<NotebookIcon />} link="/gallery"/>
                            <IconLink icon={<ClockIcon />} link="/test"/>
                            <IconLink icon={<NoteIcon />} link="/"/>
                            <IconLink icon={<StatsIcon />} link="/stats"/>
                        </div>}
                        {hover && <IconLink icon={<ProfileIcon/>} link="/profile"/>}
                    </div>
                </div>
                {hover && <SVGLine lineHeight={height} horizontal={false} strokeWidth={1}/>}
            </div>
            <div className="pt-7 pr-9 pointer-events-auto">
                <IconLink icon={<ProfileIcon/>} link="/profile" displayText="user_name"/>
            </div>
        </div>
    </div>)
}