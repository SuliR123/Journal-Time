'use client';
import ClockIcon from "@/public/clock.svg"
import NotebookIcon from "@/public/notebook.svg"
import NoteIcon from "@/public/note.svg"
import StatsIcon from "@/public/stats.svg"
import ProfileIcon from "@/public/profile.svg"
import IconLink from "./icon_link";

export default function NavBar() {
    return (
    <div className="flex flex-col items-center">
        <div className="h-[1vh]"></div>
        <div className="flex flex-row items-center justify-between gap-4 w-[85vw] h-[10vh]">
            <span className="font-hack text-[30px] text-icon-color">JournalTime</span>
            <div className="flex flex-row items-center gap-6 w-full h-full">
                <IconLink icon={<NotebookIcon />} link="/gallery"/>
                <IconLink icon={<ClockIcon />} link="/test"/>
                <IconLink icon={<NoteIcon />} link="/"/>
                <IconLink icon={<StatsIcon />} link="/stats"/>
            </div>
            <IconLink icon={<ProfileIcon className="w-6 h-6"/>} link="/profile"/>
        </div>
    </div>)
}