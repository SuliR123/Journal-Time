
interface ScrollParams {
    vertical?: boolean
    children : React.ReactNode
}

export default function Scroll({ vertical = false, children } : ScrollParams) {
    const dir = (vertical) ? "grid-col-2 grid-flow-row overflow-y-auto gap-3 w-full h-full" : "grid-rows-2 grid-flow-col overflow-x-auto gap-3 w-full h-full"
    return (
        <div className={`
                grid ${dir}
                [&::-webkit-scrollbar]:w-2 
                [&::-webkit-scrollbar]:h-2
                [&::-webkit-scrollbar-track]:bg-bg-color
                [&::-webkit-scrollbar-thumb]:bg-text-color
                [&::-webkit-scrollbar-thumb]:rounded-full
                `}>
                {children}
        </div>
    )
}