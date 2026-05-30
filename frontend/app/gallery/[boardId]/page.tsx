
export default async function Board({params} : {params : Promise<{boardId : string}>}) {
    const { boardId } = await params
    
    return (
        <div>
            {"VIEWING BOARD: " + boardId}
        </div>
    );
}