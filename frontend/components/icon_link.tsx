import Link from "next/link";

interface IconLinkProps {
    icon: React.ReactNode
    link: string
}

export default function IconLink({ icon, link } : IconLinkProps) {
    return (<Link className="w-5 h-5" href={link}>
        <span className="w-5 h-5 fill-icon-color hover:fill-text-color">{icon}</span>
        </Link>);
}