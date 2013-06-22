$(function(){
  var api_key = "474d45c709390dbb3c00102c7290c8d6fab55cd4f6b416b09"
    , api_url = "http://api.wordnik.com/v4/words.json/randomWord?";

  window.new_word = function (e){
    $.get(api_url, {
      minCorpusCount: 1000000, 
      includePartOfSpeech: "noun,verb",
      hasDictionaryDef: true,
      maxLength: 12,
      api_key: api_key
    }, function(data){
      $("[data-word]").text(data.word).val(data.word)
    })
    if(e){
      e.stopPropagation();
      e.preventDefault();
    }
    return false
  }
  if($("[data-word]").val() == "" && $("[data-word]").text() == "" ){
    window.new_word()
  }
})
