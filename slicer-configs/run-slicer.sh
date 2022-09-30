
CONFIG_PATH=$(pwd)
CONFIG_PATH="$CONFIG_PATH/config_pla_01mm_fine.ini"

prusa-slicer   \
    -g model.stl \
    --load $CONFIG_PATH \
    --scale-to-fit 30,30,30 \
    --thumbnails 400x300 \
    --output ./model.gcode

