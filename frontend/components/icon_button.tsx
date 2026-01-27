
interface IconButtonProps {
    icon: React.ReactNode;
    size?: number
    onClick?: () => void; 
    displayText?: string
}

export default function IconButton({icon, onClick = () => {alert(displayText + " clicked!")}, displayText = ""} : IconButtonProps) {
    return (
        <button className="flex flex-row justify-center items-center text-icon-color gap-1 fill-icon-color hover:fill-text-color hover:text-text-color font-hack text-[12px]" onClick={onClick}>
            <span className={`w-6 h-6`}>{icon}</span>
            <span>{displayText}</span>
        </button>
    );
}