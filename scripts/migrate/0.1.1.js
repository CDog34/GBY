// showOnIndex
function updateShowOnIndex() {
    db.article.update({}, {$set: {showonindex: false}}, {multi: true});
}
updateShowOnIndex();