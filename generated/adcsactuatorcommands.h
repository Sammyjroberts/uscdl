/**
* ADCSActuatorCommands
* Commands for the attitude control actuators
*/

#ifndef ADCSACTUATORCOMMANDS_H
#define ADCSACTUATORCOMMANDS_H

#include <stdint.h>
  #include <stdbool.h>

    /**
    * Commands for the attitude control actuators
    */
    typedef struct {
    /* Commanded reaction wheel speeds (rpm) */
    int16_t reactionWheelSpeeds[4];
    /* Commanded magnetorquer dipole moments (mA·m²) */
    int16_t magnetorquerCommands[3];
    /* Timestamp of the actuator commands (ms) */
    uint32_t commandTimestamp;
    /* Current ADCS control mode (enum) */
    uint8_t controlMode;
    } ADCSActuatorCommands_t;

    /**
    * Initialize a ADCSActuatorCommands structure with default values
    * @param p_data Pointer to the structure to initialize
    */
    void adcs_actuator_commands_init(ADCSActuatorCommands_t* p_data);

    #endif /* ADCSACTUATORCOMMANDS_H */
    