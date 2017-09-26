<?php

namespace App;

use App\User;
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
        'systolic', 'diastolic', 'pulse', 'reading_date', 'user_id',
    ];

    /**
     * Calculate the severity of this reading.
     *
     * @return array
     */
    public function severity()
    {
        if ($this->systolic <= 120 && $this->diastolic <= 80) {
            return [
                'text' => 'Normal',
                'class' => 'primary'
            ];
        } elseif (($this->systolic > 120 && $this->systolic <= 139) || ($this->diastolic > 80 && $this->diastolic <= 89)) {
            return [
                'text' => 'Pre Hypertension',
                'class' => 'warning'
            ];
        } elseif (($this->systolic >= 140 && $this->systolic <= 159) || ($this->diastolic >= 90 && $this->diastolic <= 99)) {
            return [
                'text' => 'Stage 1 Hypertension',
                'class' => 'danger'
            ];
        } elseif ($this->systolic >= 160 && $this->diastolic >= 100) {
            return [
                'text' => 'Stage 2 Hypertension',
                'class' => 'danger'
            ];
        } else {
            return [
                'text' => 'N/A',
                'class' => 'normal'
            ];
        }
    }

    /**
     * The user for whom this blood pressure belongs to.
     *
     * @return HasOne
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }

}
