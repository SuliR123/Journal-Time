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
    const [hover, setHover] = useState(false)

    return (
        <div className="flex flex-row items-start justify-between w-[7%] h-full">
            <div className="flex flex-row h-full"
                onMouseEnter={() => setHover(true)}
                onMouseLeave={() => setHover(false)}
                >
                <div className="flex flex-col items-center pl-3 w-full h-full">
                    <div className="flex flex-col items-start justify-between gap-10 w-full h-[95%] pt-7">
                        <IconLink icon={<Logo />} link="/" size="lg"/>
                        <div className="flex flex-col justify-start items-start gap-10 w-full h-full">
                            <IconLink icon={<NotebookIcon />} link="/gallery" displayText="Gallery"/>
                            <IconLink icon={<ClockIcon />} link="/test" displayText="Test"/>
                            <IconLink icon={<NoteIcon />} link="/" displayText="Create"/>
                            <IconLink icon={<StatsIcon />} link="/stats" displayText="Stats"/>
                        </div>
                        <IconLink icon={<ProfileIcon/>} link="/profile"/>
                    </div>
                </div>
            </div>
        </div>)
}