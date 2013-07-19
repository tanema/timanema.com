(function($){

  $.fn.signature = function(data){
    return $(this).each(function(){
      var canvas_el = $(this),
          canvas = canvas_el.get(0).getContext('2d');
      canvas.lineWidth = 5
      canvas.lineJoin = "round"
      canvas.lineCap = "round"
      canvas.strokeStyle = 'black';

      canvas_el.get(0).addEventListener('touchstart', _start(true));
      canvas_el.get(0).addEventListener('touchend', _end(true));
      canvas_el.get(0).addEventListener('touchcancel', _end(true));
      canvas_el.get(0).addEventListener('touchmove', _move(true));
      canvas_el.on('mousedown', _start())
      canvas_el.on('mouseup', _end())
      canvas_el.on('mousemove', _move())

      if(data && data != ""){
        console.log(data)
        var image = new Image();
        image.src = data; 
        image.onload = function() {
          canvas.drawImage(image, 0, 0);
        };
      }

      canvas_el.on("clear", function(){
        canvas_el.get(0).getContext('2d').clearRect(0,0, canvas_el.attr('width'), canvas_el.attr('height'))
        save(true)
      })

      canvas_el.on("change_color", function(event, color){
        canvas.strokeStyle = color;
      })
      
      function save(clear){
        data = clear ? "" : canvas_el.get(0).toDataURL('image/png')
        $("[data-behaviour='canvas_save']").val(data);
      }

      var is_drawing = false;
      function _start(is_touch){
        return function(event){
          is_drawing = true
          canvas.beginPath()
          canvas.moveTo(get_x(event, is_touch), get_y(event, is_touch))
          canvas.lineTo(get_x(event, is_touch)+1, get_y(event, is_touch)+1)
          canvas.stroke();
          return stopevent(event);
        }
      }
      function _end(is_touch){
        return function(event){
          is_drawing = false
          if(is_touch){
            canvas_el.touchmove();
          }else{
            canvas_el.mousemove();
          }
          canvas.closePath();
          save();
          return stopevent(event);
        }
      }
      function _move(is_touch){
        return function(event){
          if(is_drawing){
            canvas.lineTo(get_x(event, is_touch), get_y(event, is_touch))
            canvas.stroke();
            save();//hack for mobile browsers since they dont like touch end
            return stopevent(event);
          }
        }
      }
      function get_x(event, is_touch){
        return (is_touch ? event.targetTouches[0].pageX : event.clientX) / 1 - (canvas_el.offset().left - $(window).scrollLeft()) / 1;
      }
      function get_y(event, is_touch){
        return (is_touch ? event.targetTouches[0].pageY : event.clientY) / 1 - (canvas_el.offset().top - $(window).scrollTop()) / 1;
      }
      function stopevent(e){
        e.stopPropagation();
        e.preventDefault();
        return false;
      }
    })
  };

})(jQuery);
