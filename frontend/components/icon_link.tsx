import Link from "next/link";

interface IconLinkProps {
    icon: React.ReactNode
    link: string
    size?: number
    displayText?: string
}

export default function IconLink({ icon, link, displayText} : IconLinkProps) {
    return (
    <Link className="flex items-center justify-center text-icon-color gap-1 fill-icon-color hover:fill-text-color hover:text-text-color font-hack text-[12px]" href={link}>
        <span className={`w-5 h-5 fill-icon-color hover:fill-text-color flex-none`}>{icon}</span>
        <span>{displayText}</span>
    </Link> // TODO: When hovering over the user_name text the icon doesn't highlight 
    );
}