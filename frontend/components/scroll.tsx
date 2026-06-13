
interface ScrollParams {
    vertical?: boolean
    children : React.ReactNode
}

export default function Scroll({ vertical = false, children } : ScrollParams) {

    /*
    There should be: 
    a minimum number of columns (2)
    size of items in the grid should increase in size until another item could be fit into the row (max size for items essentially)
    number of columns changes dynamically
    */

    const dir = (vertical) ? "grid-cols-2 md:grid-cols-4 grid-flow-row overflow-y-auto gap-5 pr-2" : "grid-rows-2 grid-flow-col overflow-x-auto gap-5"
    return (
        <div className={`
                grid ${dir} w-full h-full
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