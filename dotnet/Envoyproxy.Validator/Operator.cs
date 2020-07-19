using System;
using Google.Protobuf.WellKnownTypes;

namespace Envoyproxy.Validator
{
    public static class Operator
    {
        #region From Protobuf Implementation

        private const int MinNanos = -Duration.NanosecondsPerSecond + 1;
        private const int MaxNanos = Duration.NanosecondsPerSecond - 1;

        private static bool IsNormalized(Duration value)
        {
            // Simple boundaries
            if (value.Nanos < MinNanos || value.Nanos > MaxNanos)
            {
                return false;
            }

            // We only have a problem if one is strictly negative and the other is
            // strictly positive.
            return Math.Sign(value.Seconds) * Math.Sign(value.Nanos) != -1;
        }

        #endregion

        public static bool Equal(Duration lhs, Duration rhs)
        {
            if (!IsNormalized(lhs))
                throw new ArgumentException("Duration is not normalized", nameof(lhs));
            if (!IsNormalized(rhs))
                throw new ArgumentException("Duration is not normalized", nameof(rhs));

            return lhs.Seconds == rhs.Seconds && lhs.Nanos == rhs.Nanos;
        }

        public static bool Lt(Duration lhs, Duration rhs)
        {
            if (!IsNormalized(lhs))
                throw new ArgumentException("Duration is not normalized", nameof(lhs));
            if (!IsNormalized(rhs))
                throw new ArgumentException("Duration is not normalized", nameof(rhs));

            return lhs.Seconds < rhs.Seconds || (lhs.Seconds == rhs.Seconds && lhs.Nanos < rhs.Nanos);
        }

        public static bool Lte(Duration lhs, Duration rhs)
        {
            if (!IsNormalized(lhs))
                throw new ArgumentException("Duration is not normalized", nameof(lhs));
            if (!IsNormalized(rhs))
                throw new ArgumentException("Duration is not normalized", nameof(rhs));

            return lhs.Seconds < rhs.Seconds || (lhs.Seconds == rhs.Seconds && lhs.Nanos <= rhs.Nanos);
        }

        public static bool Gt(Duration lhs, Duration rhs)
        {
            if (!IsNormalized(lhs))
                throw new ArgumentException("Duration is not normalized", nameof(lhs));
            if (!IsNormalized(rhs))
                throw new ArgumentException("Duration is not normalized", nameof(rhs));

            return lhs.Seconds > rhs.Seconds || (lhs.Seconds == rhs.Seconds && lhs.Nanos > rhs.Nanos);
        }

        public static bool Gte(Duration lhs, Duration rhs)
        {
            if (!IsNormalized(lhs))
                throw new ArgumentException("Duration is not normalized", nameof(lhs));
            if (!IsNormalized(rhs))
                throw new ArgumentException("Duration is not normalized", nameof(rhs));

            return lhs.Seconds > rhs.Seconds || (lhs.Seconds == rhs.Seconds && lhs.Nanos >= rhs.Nanos);
        }

        #region From Protobuf Implementation

        private const long BclSecondsAtUnixEpoch = 62135596800;

        // Duplicated from Protobuf implementation because Protobuf's Timestamp.FromDateTime()
        // expects UTC and the timestamp validation rule is targeting local time.
        public static Timestamp Now()
        {
            var now = DateTime.Now;

            long secondsSinceBclEpoch = now.Ticks / TimeSpan.TicksPerSecond;
            int nanoseconds = (int) (now.Ticks % TimeSpan.TicksPerSecond) * Duration.NanosecondsPerTick;

            return new Timestamp
            {
                Seconds = secondsSinceBclEpoch - BclSecondsAtUnixEpoch,
                Nanos = nanoseconds
            };
        }

        #endregion
    }
}
