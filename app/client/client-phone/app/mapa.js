let AttachDragTo = (function() {
    let _AttachDragTo = function(el) {
      this.el = el;
      this.mouse_is_down = false;
  
      this.init();
    };
  
    _AttachDragTo.prototype = {
      onMousemove: function(e) {
        if (!this.mouse_is_down) return;
  
        let tg = e.target,
          x = e.clientX,
          y = e.clientY,
          target_width = tg.clientWidth,
          target_height = tg.clientHeight,
          image_proportion,
          image_height = 600, //Change this variable when changing the image.
          image_width = 800, //Change this variable when changing the image.
          max_pos_x = 0,
          max_pos_y = 0;
        
        image_proportion = image_width / image_height;
        
        if(image_width > target_width && image_height > target_height) {
            max_pos_y = target_width / image_proportion - target_height;
          } else {
            if (target_width - image_width  > target_height - image_height) {
              max_pos_y = target_width / image_proportion - target_height;
            } else {
              max_pos_x = target_width / image_proportion - target_width;
            }
          }
  
          let image_bg_pos_x = x - this.origin_x + this.origin_bg_pos_x;
          let image_bg_pos_y = y - this.origin_y + this.origin_bg_pos_y;
  
          if (image_bg_pos_x < 0 && image_bg_pos_x > -max_pos_x) {
            tg.style.backgroundPositionX = image_bg_pos_x + 'px';
          }
  
          if (image_bg_pos_y < 0 && image_bg_pos_y > -max_pos_y) {
            tg.style.backgroundPositionY = image_bg_pos_y + 'px';
          }
      },
  
      onMouseleave: function(e) {
        this.mouse_is_down = false;
        
        let tg = e.target,
          styles = getComputedStyle(tg);
  
        this.origin_bg_pos_x = parseInt(
          styles.getPropertyValue('background-position-x'),
          10
        );
        this.origin_bg_pos_y = parseInt(
          styles.getPropertyValue('background-position-y'),
          10
        );
        
        tg.style.cursor = 'grab';
      },
  
      onMousedown: function(e) {
        e.target.style.cursor = 'grabbing';
        
        this.mouse_is_down = true;
        this.origin_x = e.clientX;
        this.origin_y = e.clientY;
      },
  
      onMouseup: function(e) {
        let tg = e.target,
          styles = getComputedStyle(tg);
  
        this.mouse_is_down = false;
        
        this.origin_bg_pos_x = parseInt(
          styles.getPropertyValue('background-position-x'),
          10
        );
        this.origin_bg_pos_y = parseInt(
          styles.getPropertyValue('background-position-y'),
          10
        );
        
        tg.style.cursor = 'grab';
      },
  
      init: function() {
        let styles = getComputedStyle(this.el);
        this.origin_bg_pos_x = parseInt(
          styles.getPropertyValue('background-position-x'),
          10
        );
        this.origin_bg_pos_y = parseInt(
          styles.getPropertyValue('background-position-y'),
          10
        );
        
        let imageUrl = this.el.style.backgroundImage.replace(/url\((['"])?(.*?)\1\)/gi, '$2');
        let image = new Image();
        image.src = imageUrl;
  
        image.onload = () => {
          this.image_width = image.width,
            this.image_height = image.height;
        }
  
  
        //attach events
        this.el.addEventListener('mousedown', this.onMousedown.bind(this), false);
        this.el.addEventListener('mouseup', this.onMouseup.bind(this), false);
        this.el.addEventListener('mousemove', this.onMousemove.bind(this), false);
        this.el.addEventListener(
          'mouseleave',
          this.onMouseleave.bind(this),
          false
        );
      }
    };
  
    return function(el) {
      new _AttachDragTo(el);
    };
  })();
  
  /*** IMPLEMENTATION ***/
  //1. Get your element.
  const image = document.getElementById('image');
  //2. Attach the drag.
  AttachDragTo(image);