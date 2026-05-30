export default async function ViewNote({ params } : {params: Promise<{noteId : string}>}) {
    const { noteId } = await params
    // Display the Note normally first (scrollable with no outline ),
    // have an edit button that allows the user to edit the textbox, render buttons to edit
    return (
        <div>
            {"NOTE: " + noteId}
        </div>
    );
}