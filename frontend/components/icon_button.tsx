
interface IconButtonProps {
    icon: React.ReactNode;
    onClick: () => void; 
}

export default function IconButton({icon, onClick} : IconButtonProps) {
    return (
        <button className="w-6 h-6" onClick={onClick}>
            <span className="w-6 h-6 fill-icon-color hover:fill-text-color">{icon}</span>
        </button>
    );
}