/**
* ADCSAttitudeState
* Current attitude state of the spacecraft
*/

#ifndef ADCSATTITUDESTATE_H
#define ADCSATTITUDESTATE_H

#include <stdint.h>
  #include <stdbool.h>

    /**
    * Current attitude state of the spacecraft
    */
    typedef struct {
    /* Quaternion representing the spacecraft attitude (quaternion) */
    float quaternion[4];
    /* Angular velocity vector of the spacecraft (rad/s) */
    float angularVelocity[3];
    /* Timestamp of the attitude measurement (ms) */
    uint32_t timestamp;
    /* Current mode of attitude determination (enum) */
    uint8_t attitudeDeterminationMode;
    /* Flag indicating if the attitude solution is valid */
    bool attitudeValid;
    } ADCSAttitudeState_t;

    /**
    * Initialize a ADCSAttitudeState structure with default values
    * @param p_data Pointer to the structure to initialize
    */
    void adcs_attitude_state_init(ADCSAttitudeState_t* p_data);

    #endif /* ADCSATTITUDESTATE_H */
    