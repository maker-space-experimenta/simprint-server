
CONFIG_PATH=$(pwd)
CONFIG_PATH="$CONFIG_PATH/slicer-configs/config_pla_03mm_draft.ini"

prusa-slicer   \
    -g model.stl \
    --load $CONFIG_PATH \
    --scale-to-fit 30,30,30 \
    --thumbnails 400x300 \
    --output _temp/model.gcode

