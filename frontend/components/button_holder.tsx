interface ButtonHolderParams {
    transparent?: boolean
    children : React.ReactNode
}


export default function ButtonHolder({ transparent = false, children } : ButtonHolderParams) {
    const backdrop = transparent ? '' : 'backdrop-brightness-125'
    
    return (
    <div className="bg-background-color">
        <div className={`flex flex-row w-[15vw] h-[6vh] justify-around items-center rounded-lg ${backdrop}`}> 
            {children}
        </div>
    </div>
    )
}