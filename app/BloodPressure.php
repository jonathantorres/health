<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class BloodPressure extends Model
{
    /**
     * The table associated with the model.
     *
     * @var string
     */
    protected $table = 'blood_pressures';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'systolic', 'diastolic', 'pulse', 'reading_date',
    ];

    /**
     * Calculate the severity of this reading.
     *
     * @return string
     */
    public function severity()
    {
        if ($this->systolic <= 120 && $this->diastolic <= 80) {
            return 'Normal';
        } elseif (($this->systolic > 120 && $this->systolic <= 139) || ($this->diastolic > 80 && $this->diastolic <= 89)) {
            return 'Pre Hypertension';
        } elseif (($this->systolic >= 140 && $this->systolic <= 159) || ($this->diastolic >= 90 && $this->diastolic <= 99)) {
            return 'Stage 1 Hypertension';
        } elseif ($this->systolic >= 160 && $this->diastolic >= 100) {
            return 'Stage 2 Hypertension';
        } else {
            return 'N/A';
        }
    }

}
