
export default async function BoardPage({params} : {params : Promise<{boardId : string}>}) {
    const { boardId } = await params
    
    return (
        <div>
            {"VIEWING BOARD: " + boardId}
        </div>
    );
}